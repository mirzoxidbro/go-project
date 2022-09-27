<?php

namespace App\Http\Controllers;

use App\Http\Requests\OrderItemRequest;
use App\Http\Services\OrderItemService;
use Illuminate\Http\Request;

class OrderItemController extends Controller
{
    protected $service;

    public function __construct()
    {
        $this->service = new OrderItemService;
    }

    public function index(){
        $orderitem = $this->service->index();
        return response()->json([
            'orderitem' => $orderitem
        ]);
    }

    public function store(OrderItemRequest $request){
        $orderitem = $this->service->store($request->validated());
        return response()->json([
            'orderitem' => $orderitem
        ]);
    }

    public function show($id){
        $orderitem = $this->service->show($id);
        return response()->json([
            'orderitem' => $orderitem
        ]);
    }

    public function update(OrderItemRequest $request, $id){
        $orderitem = $this->service->update($request->validated(), $id);
        return response()->json([
            'orderitem' => $orderitem
        ]);
    }

    public function destroy($id){
     $orderitem = $this->service->destroy($id);
     return response()->json([
            'message' => 'deleted successfully'
        ]);
    }

}
